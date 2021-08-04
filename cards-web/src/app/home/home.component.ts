import { Component, OnInit } from '@angular/core';
import { Game, GameApiService, GameResponse } from '../game-api.service';
import { FormBuilder, FormGroup, FormArray, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { GameService } from '../game.service';


@Component({
  selector: 'cards-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  gameData!: Game;

  gameForm!: FormGroup;

  constructor(
    private gameApiService: GameApiService,
    private fb: FormBuilder,
    private router: Router,
    private gameService: GameService,
  ) { }

  ngOnInit(): void {
    this.gameForm = this.fb.group({
      numDecks: [null,[Validators.required]],
      players: this.fb.array([this.createPlayer()], Validators.required)
    });
  }

  createPlayer(): FormGroup {
    return this.fb.group({
      name: [null, Validators.required],
    });
  }

  addPlayer(): void {
    this.players.push(this.createPlayer());
  }

  removerPlayer(index: number): void {
    if (this.players.length > index && this.players.length > 1) {
      this.players.removeAt(index);
    }
  }

  get players(): FormArray {
    return this.gameForm.get('players') as FormArray;
  }

  public createGame() {
    let numDecks = parseInt(this.gameForm.get('numDecks')?.value);
    let playerNames = this.gameForm.get('players')?.value.map((x: {name: string}) => x.name); // [{name: (string)}, ...]

    this.gameApiService.createGame(numDecks, playerNames).subscribe((response: GameResponse) => {
      this.gameService.game = response.data;
      this.router.navigate(['game', this.gameService.key]);
    });
  }
}
