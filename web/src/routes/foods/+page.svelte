<script lang="ts">
	import IconPlus from '@tabler/icons-svelte/IconPlus.svelte';
	import IconMinus from '@tabler/icons-svelte/IconMinus.svelte';
	import type { PageServerData } from './$types';
	import type { Food } from '$lib/models/Food';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { writable } from 'svelte/store';
	import type { WithData, APIResponse } from '$lib/models/Responses';

	export let data: PageServerData;
	const foods = writable<Food[]>(data.foods);

	async function changeQuantity(action: 'increment' | 'decrement', food: Food) {
		const response = await fetch(`${PUBLIC_API_URL}/api/v1/foods/${food.id}`, {
			method: 'PUT',
			body: JSON.stringify({
				name: food.name,
				quantity: action === 'decrement' ? Math.max(0, food.quantity - 1) : food.quantity + 1
			}),
			headers: {
				accept: 'application/json',
				'content-type': 'application/json'
			}
		});
		const result: WithData<APIResponse, Food> = await response.json();
		$foods = $foods.map((food) => (food.id !== result.data.id ? food : result.data));
	}
</script>

<h1>Foods here</h1>
<div class="relative w-full grid grid-cols-2 gap-5 p-5">
	{#each $foods as food}
		<div class="relative w-full aspect-square rounded border border-r-red-200">
			<h2>{food.name}</h2>
			<span>{food.quantity} left in stock</span>
			<div class="relative flex justify-center gap-5">
				<button on:click={() => changeQuantity('decrement', food)}>
					<IconMinus />
				</button>
				<button on:click={() => changeQuantity('increment', food)}>
					<IconPlus />
				</button>
			</div>
		</div>
	{/each}
</div>
