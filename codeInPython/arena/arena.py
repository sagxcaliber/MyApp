from codeInPython.models.models import Player


def new_player(*, name: str, health: int, strength: int, attack: int):
    """

    :param name: need to add the name here to set for the player as string
    :param health: need to add the health as int
    :param strength: need to add strength as int
    :param attack:  need to attack to add as int
    :return:  Player dict objects {'name': 'Deepak', 'health': 50, 'strength': 3, 'attack': 20}
    """
    if health <= 0 | strength <= 0 | attack <= 0:
        # validation checks added
        raise Exception("All Attributes must be positive integers")
    else:
        # creating player object
        create_player = Player()
        create_player.set(name, health, strength, attack)
        return create_player.get()
