# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: game.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import game_pb2


class GameServiceBase(abc.ABC):

    @abc.abstractmethod
    async def Join(self, stream: 'grpclib.server.Stream[game_pb2.NewPlayer, game_pb2.PlayerID]') -> None:
        pass

    @abc.abstractmethod
    async def InitialState(self, stream: 'grpclib.server.Stream[game_pb2.NewPlayerInitialState, game_pb2.PlayerReady]') -> None:
        pass

    @abc.abstractmethod
    async def Turn(self, stream: 'grpclib.server.Stream[game_pb2.NewTurn, game_pb2.NewAction]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/GameService/Join': grpclib.const.Handler(
                self.Join,
                grpclib.const.Cardinality.UNARY_UNARY,
                game_pb2.NewPlayer,
                game_pb2.PlayerID,
            ),
            '/GameService/InitialState': grpclib.const.Handler(
                self.InitialState,
                grpclib.const.Cardinality.UNARY_UNARY,
                game_pb2.NewPlayerInitialState,
                game_pb2.PlayerReady,
            ),
            '/GameService/Turn': grpclib.const.Handler(
                self.Turn,
                grpclib.const.Cardinality.UNARY_UNARY,
                game_pb2.NewTurn,
                game_pb2.NewAction,
            ),
        }


class GameServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.Join = grpclib.client.UnaryUnaryMethod(
            channel,
            '/GameService/Join',
            game_pb2.NewPlayer,
            game_pb2.PlayerID,
        )
        self.InitialState = grpclib.client.UnaryUnaryMethod(
            channel,
            '/GameService/InitialState',
            game_pb2.NewPlayerInitialState,
            game_pb2.PlayerReady,
        )
        self.Turn = grpclib.client.UnaryUnaryMethod(
            channel,
            '/GameService/Turn',
            game_pb2.NewTurn,
            game_pb2.NewAction,
        )
