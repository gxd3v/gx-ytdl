import {constants} from './constants';

export function SaveTile(tileID: number): void {
  localStorage.setItem(constants.storedTileName, tileID.toString());
}

export function RetrieveTile(): number | null {
  const tile: string | null = localStorage.getItem(constants.storedTileName);
  if (tile !== null) {
    return parseInt(tile);
  }

  return null;
}
