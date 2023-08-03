class board:
    __myboard=None
    def __init__(self,name):
        self.name=name 
    @staticmethod
    def createBoard(name):
        if __myboard==None:
            __myboard=board(name)
            return __myboard
        else:
            return __myboard

if __name__=="__main__":
    newboard=board.createBoard("yash")
    newboard2=board.createBoard("shah")
    print(newboard.name)
    print(newboard2.name)