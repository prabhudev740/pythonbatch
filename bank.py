# Customer 
# ---------
# Branch name
# PIN Code
# Account no.
# First Name
# Last Name
# Amount
# ------------
# Deposite
# Withdraw
# Display --> Show the ammount


class CitiBank:
    branch_name = 'Bangalore'
    pin_code = 560076
    base_number = 110000

    def __init__(self, fn, ln, amount=0):
        self.first_name = fn
        self.last_name = ln
        self.balance = amount
        self.accout_no = CitiBank.increment()

    @classmethod
    def increment(cls):
        cls.base_number = cls.base_number + 1
        return cls.base_number

    def details(self):
        print(f"Branc Name: {self.branch_name}")
        print(f"PIN Code: {self.pin_code}")
        print(f"Account no: {self.accout_no}")
        print(f"First Name: {self.first_name}")
        print(f"Last Name: {self.last_name}")
        print(f"Ammount: {self.balance}")
        print()

    def deposite(self, amt):
        self.balance = self.balance + amt
        print("Successfully Deposited")

    def withdraw(self, amt):
        if self.balance >= amt:
            self.balance -= amt
            print('Successfully withdrawn')
        else:
            print('you dont have sufficient bal')

    def display(self):
        print(f"Curent Bal: {self.balance}")


prabhu = CitiBank("Prabhu", 'P', 500)
prabhu.details()


prabhu.display()
prabhu.deposite(500)
prabhu.display()
prabhu.withdraw(200)
prabhu.display()
prabhu.withdraw(400)
prabhu.display()
prabhu.withdraw(500)
prabhu.display()



# print()

# suresh = CitiBank('Suresh', 'K')
# suresh.details()