#Abstract Class
class Animal():

    def __init__(self, name):
        self.name = name
        print("Hello, I'm {}".format(self.name))

    def who_ami_i(self):
        raise NotImplementedError("Cannot create instances of an abstract class")

class Dog(Animal):

    def who_am_i(self):
        print("I am a dog")

    def speak(self):
        return self.name + " says Woof!"

class Cat(Animal):
    
    def speak(self):
        return self.name + " says Meow!"

niko = Dog("niko")
felix = Cat("felix")

#Polymorphism

def pet_speak(pet):
    print(pet.speak())

pet_speak(niko)
pet_speak(felix)