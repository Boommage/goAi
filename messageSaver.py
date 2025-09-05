import asyncio
import sqlite3
import time
from winrt.windows.ui.notifications.management import UserNotificationListener, UserNotificationListenerAccessStatus

DB_FILE = "messages.db"

def init_db():
    conn = sqlite3.connect(DB_FILE)
    conn.execute("PRAGMA journal_mode=WAL;")
    c = conn.cursor()
    c.execute("""CREATE TABLE IF NOT EXISTS messages (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    username TEXT,
                    message TEXT,
                    processed INTEGER DEFAULT 0
                )""")
    conn.commit()
    conn.close()

def save_message(username, message):
    conn = sqlite3.connect(DB_FILE)
    conn.execute("PRAGMA journal_mode=WAL;")
    c = conn.cursor()
    c.execute("INSERT INTO messages (username, message) VALUES (?, ?)", (username, message))
    conn.commit()
    conn.close()  

def message_exists(username, message):
    conn = sqlite3.connect(DB_FILE)
    conn.execute("PRAGMA journal_mode=WAL;")
    c = conn.cursor()
    c.execute("SELECT 1 FROM messages WHERE username=? AND message=? LIMIT 1", (username, message))
    exists = c.fetchone() is not None
    conn.close()
    return exists

async def main():
    init_db()
    listener = UserNotificationListener.get_current()
    status = await listener.request_access_async()

    while True:
        if status == UserNotificationListenerAccessStatus.ALLOWED:
            notifs = await listener.get_notifications_async(0)
            for notif in notifs:
                app_info = notif.app_info.display_info.display_name

                if app_info.lower() == "discord":
                    tempUsername = notif.notification.visual.bindings[0].get_text_elements()[0].text
                    tempUserMessage = notif.notification.visual.bindings[0].get_text_elements()[1].text

                    if not message_exists(tempUsername, tempUserMessage):
                        save_message(tempUsername, tempUserMessage)
                        print("Saved:", tempUsername, ":", tempUserMessage)
                        time.sleep(5)
                        


asyncio.run(main())