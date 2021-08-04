import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';
import { Card, Game, Player } from '../game-api.service';
import { GameService } from '../game.service';

@Component({
  selector: 'cards-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.scss']
})
export class GameComponent implements OnInit {
  playerUUID: string = "";
  shuffling: boolean = false;

  constructor(
    private route: ActivatedRoute,
    private gameService: GameService,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      let key = params.get('key');
      if (key !== null && this.gameService.key != key) {
        this.gameService.loadGame(key).subscribe((loaded: boolean) => {
          if (!loaded) {
            this.router.navigate(['home']);
          }
        });
      }
    });
  }

  get game(): Game|null {
    return this.gameService.game;
  }

  get player(): Player|null {
    if (this.playerUUID != '' && this.game) {
      return this.game.Players.find((player: Player) => player.UUID == this.playerUUID) ?? null;
    }
    return null;
  }

  get lastDiscard(): Card|null {
    let discardCards = this.game?.Discard;
    if (discardCards && discardCards.length > 0) {
      return discardCards[discardCards.length - 1];
    }
    return null;
  }

  public shuffle(): void {
    this.shuffling = true;
    this.gameService.shuffle().subscribe((shuffled: boolean) => {
      if (shuffled) {
        this.shuffling = false;
      } else {
        this.router.navigate(['home']);
      }
    });
  }

  public shuffleDiscard(): void {
    alert("Not implemnted");
  }

  public dealCard(): void {
    this.gameService.dealCards(1, this.playerUUID);
  }

  public discard(uuid: string): void {
    this.gameService.discard(uuid, this.playerUUID);
  }
}

