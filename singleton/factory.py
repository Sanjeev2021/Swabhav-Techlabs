class TableFan:
    def __init__(self):
        pass

    def print():
        print("TableFan12S")
class EFan:
    def __init__(self):
        pass

    def print():
        print("EFan")
class FanFactory:
    def createFan(self,type):
        switcher={
            "TableFan":TableFan,
            "EFan":EFan
        }
        return switcher.get(type)

if __name__ == "__main__":
    f=FanFactory()
    t=f.createFan("TableFan")
    e=f.createFan("EFan")
    t.print()
    e.print()

