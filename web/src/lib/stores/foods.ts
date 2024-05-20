import type { Food } from "$models/Food";
import { writable } from "svelte/store";

export const foods = writable<Food[]>([]);
