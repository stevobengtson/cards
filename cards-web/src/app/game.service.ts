import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { map } from 'rxjs/operators'
import { Game, GameApiService, GameResponse } from './game-api.service';

@Injectable({
  providedIn: 'root'
})
export class GameService {
  private currentGame!: Game|null;

  constructor(private gameApiService: GameApiService) { }

  set game(val: Game|null) {
    this.currentGame = val;
  }

  get game(): Game|null {
    return this.currentGame;
  }

  get key(): string|null {
    return this.currentGame?.Key ?? null;
  }

  public loadGame(key: string): Observable<boolean> {
    if (this.currentGame && key == this.currentGame.Key) {
      return new Observable<boolean>((observer) => {
        observer.next(true);
      });
    }

    return this.gameApiService.getGame(key).pipe(
      map((response: GameResponse) => {
        if (response && response.data && response.data.Key != "") {
          this.currentGame = response.data;
        } else {
          this.currentGame = null;
        }
        return this.currentGame !== null;
      })
    );
  }

  public shuffle(): Observable<boolean> {
    if (!this.currentGame || this.currentGame.Key == "") {
      return new Observable<boolean>((observer) => {
        observer.next(false);
      });
    }

    return this.gameApiService.shuffle(this.currentGame.Key).pipe(
      map((response: GameResponse) => { return true; })
    );
  }

  public dealCards(count: number, playerUUID: string): void {
    let key = this.key;
    if (key != null) {
      this.gameApiService.deal(key, count, playerUUID).subscribe((response: GameResponse) => {
        this.game = response.data;
      });
    }
  }

  public discard(cardUUID: string, playerUUID: string): void {
    let key = this.key;
    if (key != null) {
      this.gameApiService.discard(key, [cardUUID], playerUUID).subscribe((response: GameResponse) => {
        this.game = response.data;
      });
    }
  }
}
