class Player:
    def __init__(self):
        self.Name: str = ''
        self.Health: int = 0
        self.Strength: int = 0
        self.Attack: int = 0

    def get(self):
        return {
            "name": self.Name
            , "health": self.Health
            , "strength": self.Strength
            , "attack": self.Attack
        }

    def set(self, name, health, strength, attack):
        self.Name = name
        self.Health = health
        self.Strength = strength
        self.Attack = attack
