#!/usr/bin/python
# -*- coding: utf-8 -*-

import engine, train

from bots.randbot import RandBot
from bots.reinforce import REINFORCE


# ==============================================================================
# MAIN
# Main process for simulating matches of different types of bots
# ==============================================================================

if __name__ == "__main__":
    # Set the map
    cfg_file = "maps/grid.txt"
    
    # Set the bots to play the game
    bots = [REINFORCE(), RandBot()]

    NUM_TRAINING_EPISODES = 5
    MAX_ROUNDS = 1000

    for i in range(1, NUM_TRAINING_EPISODES+1):
        config = engine.GameConfig(cfg_file)
        game = engine.Game(config, len(bots))

        iface = train.Interface(game, bots, debug=False)
        iface.run(max_rounds=MAX_ROUNDS)

        for bot in bots:
            if bot.save_model:
                bot.save_trained_model()

       