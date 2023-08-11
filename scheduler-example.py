import schedule
import time

def job():
    print("I'm working...")

schedule.every(11).seconds.do(job)
schedule.every(1).minutes.do(job)

def job_with_argument(name):
    print(f"I am {name}")

schedule.every(12).seconds.do(job_with_argument, name="Peter")

while True:
    schedule.run_pending()
    # time.sleep(1)