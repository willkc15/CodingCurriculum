class Animal():

    def __init__(self):
        print("Animal Created")

    def who_ami_i(self):
        print("I am an animal")

    def eat(self):
        print("I am eating")

class Dog():

    def __init__(self):
        Animal.__init__(self)
        print("Dog created")