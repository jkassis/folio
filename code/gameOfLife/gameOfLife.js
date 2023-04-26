// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

/*
The universe of the Game of Life is an infinite two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, alive or dead, or "populated" or "unpopulated". Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

1) Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2) Any live cell with two or three live neighbours lives on to the next generation.
3) Any live cell with more than three live neighbours dies, as if by overpopulation.
4) Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

The initial pattern constitutes the seed of the system. The first generation is created by applying the above rules simultaneously to every cell in the seed — births and deaths occur simultaneously, and the discrete moment at which this happens is sometimes called a tick (in other words, each generation is a pure function of the preceding one). The rules continue to be applied repeatedly to create further generations.
*/

class Grid {
  constructor(seed) {
    this.grid = seed;
    this.width = seed[0].length;
    this.height = seed.length;
  }

  get (x, y) {
    return this.grid[y][x];
  }

  getNeighbors(x, y) {
    var startX  = Math.max(x - 1, 0);
    var endX = Math.min(x+1, this.width - 1);

    var startY  = Math.max(y - 1, 0);
    var endY = Math.min(y+1, this.height - 1);

    var living = 0;
    var dead = 0;
    for (var xi = startX; xi <= endX; xi++) {
      for (var yi = startY; yi <= endY; yi++) {
        if (xi == x && yi == y) {
          continue;
        }

        if (this.get(xi,yi)) {
          living++;
        } else {
          dead++;
        }
      }
    }

    return [living, dead];
  }

  print() {
    for (var row of this.grid) {
      console.log(row);
    }
    console.log();
  }

  tick() {
    var nextGrid = [];
    for (var row of this.grid) {
      nextGrid.push(row.slice(0));
    }

    for(var x = 0 ; x < this.width ; x++) {
      for(var y = 0 ; y < this.height ; y++) {
        var [living, dead] = this.getNeighbors(x, y);
        if (this.get(x,y)) {
          if (living < 2) {
            nextGrid[y][x] = 0;
          } else if (living < 4) {
            nextGrid[y][x] = 1;
          } else {
            nextGrid[y][x] = 0;
          }
        } else {
          if (living == 3) {
            nextGrid[y][x] = 1;
          }
        }
      }
    }

    this.grid = nextGrid;
  }
}

var seed = [
  [0, 1, 0, 0, 0],
  [0, 0, 1, 0, 0],
  [1, 1, 1, 0, 0],
  [0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0]
];

/*
var observed =  [
[ 0, 0, 0, 0, 0 ]
[ 1, 0, 0, 0, 0 ]
[ 0, 0, 1, 0, 0 ]
[ 0, 1, 0, 0, 0 ]
[ 0, 0, 0, 0, 0 ]
];

var expected = [
  [0, 0, 0, 0, 0],
  [1, 0, 1, 0, 0],
  [0, 1, 1, 0, 0],
  [0, 1, 0, 0, 0],
  [0, 0, 0, 0, 0]
];
*/


var grid = new Grid(seed);
grid.print();


// console.log(grid.get(0,0));
// console.log(grid.get(0,1));
// console.log(grid.get(0,2));

// console.log(grid.getNeighbors(0,0));
// console.log(grid.getNeighbors(1,0));
// console.log(grid.getNeighbors(2,0));

// console.log(grid.getNeighbors(0,1));
// console.log(grid.getNeighbors(1,1));
// console.log(grid.getNeighbors(2,1));

// console.log(grid.getNeighbors(0,2));
// console.log(grid.getNeighbors(1,2));
// console.log(grid.getNeighbors(2,2));

grid.tick();
grid.print();









