from codeInPython.arena.arena import *


def main() -> str:
    """
    Creating this function to execute the main file with the following code

    player1 need to be created first
    then will create player2
    after that we will call the fight function to get the result
    :return: we will get the result of fight
    """

    player1 = new_player(name="Sagar", health=100, strength=5, attack=10)
    player2 = new_player(name="Deepak", health=50, strength=3, attack=20)

    print(player1)
    print(player2)
    # Need to add fight logic here
    result = ''

    return result


if __name__ == "__main__":
    """
    creating main execute  condition
    """
    main()
