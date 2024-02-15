from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine, Column, Integer, String, Date
from datetime import datetime, timedelta
from sqlalchemy.orm import sessionmaker

engine = create_engine('sqlite:///todo.db?check_same_thread=False')
Base = declarative_base()


class Table(Base):
    __tablename__ = 'task'
    id = Column(Integer, primary_key=True)
    task = Column(String, default='default_value')
    deadline = Column(Date, default=datetime.today())

    def __repr__(self):
        return self.task


while True:
    Base.metadata.create_all(engine)
    Session = sessionmaker(bind=engine)
    session = Session()
    rows = session.query(Table).all()
    today = datetime.today()
    print("1) Today's tasks")
    print("2) Week's tasks")
    print('3) All tasks')
    print('4) Missed tasks')
    print('5) Add task')
    print('6) Delete task')
    print('0) Exit')
    choice = input()
    if choice == '0':
        break
    elif choice == '1':
        rows = session.query(Table).filter(Table.deadline == today.date()).all()
        print(f"Today {today.day} {today.strftime('%b')}:")
        if not rows:
            print('Nothing to do!')
        else:
            counter = 1
            for row in rows:
                print(f'{counter}. {row.task}')
                counter += 1
        print()
    elif choice == '2':
        for i in range(7):
            next_day = datetime.today() + timedelta(days=i)
            day_rows = session.query(Table).filter(Table.deadline == next_day.date()).all()
            print(f"{next_day.strftime('%A')} {next_day.day} {next_day.strftime('%b')}:")
            if not day_rows:
                print('Nothing to do!')
            else:
                counter = 1
                for row in day_rows:
                    print(f'{counter}. {row.task}')
                    counter += 1
            print()
    elif choice == '3':
        rows = session.query(Table).order_by(Table.deadline).all()
        print('All tasks:')
        if not rows:
            print('Nothing to do!')
        else:
            counter = 1
            for row in rows:
                print(f"{counter}. {row.task}. {row.deadline.day} {row.deadline.strftime('%b')}")
                counter += 1
        print()
    elif choice == '4':
        rows = session.query(Table).filter(Table.deadline < today.date()).all()
        print('Missed tasks:')
        if not rows:
            print('Nothing is missed!')
        else:
            counter = 1
            for row in rows:
                print(f"{counter}. {row.task}. {row.deadline.day} {row.deadline.strftime('%b')}")
                counter += 1
        print()
    elif choice == '5':
        try:
            print('Enter task')
            new_task = input()
            print('Enter deadline')
            new_deadline = input()
            dl = datetime.strptime(new_deadline, '%Y-%m-%d')
            new_row = Table(task=new_task, deadline=dl)
            session.add(new_row)
            session.commit()
            print('The task has been added!')
        except ValueError:
            print('Deadline should match the format YYYY-MM-DD')
        print()
    elif choice == '6':
        rows = session.query(Table).order_by(Table.deadline).all()
        print('Choose the number of the task you want to delete:')
        if not rows:
            print('Nothing to delete!')
        else:
            counter = 1
            for row in rows:
                print(f"{counter}. {row.task}. {row.deadline.day} {row.deadline.strftime('%b')}")
                counter += 1
            del_num = int(input())
            session.delete(rows[del_num-1])
            session.commit()
        print()
