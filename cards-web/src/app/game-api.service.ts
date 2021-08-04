import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface Card {
  UUID: string;
  Type: string;
  Suit: string;
}

export interface Deck {
  Cards: Card[];
}

export interface Player {
  UUID: string;
  Name: string;
  Position: string;
  Cards: Card[];
}

export interface Game {
  UUID: string;
  Key: string;
  Deck: Deck;
  Discard: Card[];
  Players: Player[];
}

export interface GameResponse {
  data: Game;
}


@Injectable({
  providedIn: 'root'
})
export class GameApiService {
  private REST_API_SERVER = "http://localhost:8080";

  constructor(private httpClient: HttpClient) { }

  public createGame(numDecks: number, playerNames: string[]) {
    let parameters = {
      decks: numDecks,
      players: playerNames,
    };

    return this.httpClient.post<GameResponse>(`${this.REST_API_SERVER}/game`, parameters);
  }

  public getGame(key: string) {
    return this.httpClient.get<GameResponse>(`${this.REST_API_SERVER}/game/${key}`);
  }

  public shuffle(key:string) {
    return this.httpClient.put<GameResponse>(`${this.REST_API_SERVER}/game/${key}/shuffle`, {});
  }

  public deal(key: string, count: number, playerUUID: string) {
    let parameters = {
      count: count,
      player_uuid: playerUUID,
    };

    return this.httpClient.put<GameResponse>(`${this.REST_API_SERVER}/game/${key}/deal`, parameters);
  }

  public discard(key: string, cardUUIDs: string[], playerUUID: string) {
    let parameters = {
      cards: cardUUIDs,
      player_uuid: playerUUID,
    };

    return this.httpClient.put<GameResponse>(`${this.REST_API_SERVER}/game/${key}/discard`, parameters);
  }
}
