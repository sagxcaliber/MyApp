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


def turn_decide(player_a: dict, player_b: dict):
    if player_a.get('health') < player_b.get('health'):
        return player_a, player_b
    else:
        player_b, player_a

def rollDie():
    ...


def attack(attacker, defender):
    attackerRoll = ''
    defendRoll = ''


def fight(player_a: dict, player_b: dict) -> str:
    while player_a.get('health') > 0 and player_b.get('health') > 0:
        # here we need to add a turn decider function
        attacker, defender = turn_decide(player_a, player_b)
        # Now as we now the attacker and defender we need to make attacks
        attack(attacker, defender)
        if defender.get('health') <= 0:
            return "{} wins!!".format(attacker.get('name'))
        if attacker.get('health') <= 0:
            return "{} wins!!".format(defender.get('name'))

    return "Its a draw"
