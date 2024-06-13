<script lang="ts">
	import IconPlus from '@tabler/icons-svelte/IconPlus.svelte';
	import IconSearch from '@tabler/icons-svelte/IconSearch.svelte';
	import IconX from '@tabler/icons-svelte/IconX.svelte';
	import type { PageServerData } from './$types';
	import type { Food } from '$lib/models/Food';
	import AddModal from './AddModal.svelte';
	import { foods } from '$stores/foods';
	import EditModal from './EditModal.svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	export let data: PageServerData;

	$foods = data.foods;
	let show = false;
	let selectedFood: Food | null = null;

	$: search = $page.url.searchParams.get('search') ?? '';
	$: filteredFoods = $foods.filter((food) => search === '' || food.name.toLowerCase().includes(search.toLowerCase()));
</script>

<svelte:head>
	<title>Food Inventory - My Home</title>
</svelte:head>

<main class="relative w-full h-full flex flex-col px-5 pb-5 gap-5">
	<h1>Food Inventory</h1>
	<form action="/food-inventory" class="relative w-full min-h-[50px] flex bg-slate-200 rounded">
		<input type="search" name="search" value={search} placeholder="Search food" class="relative w-full h-full px-5 rounded-l bg-transparent" />
		<button
			type="button"
			class="relative h-full aspect-square bg-transparent flex justify-center items-center {search === '' ? 'invisible' : ''}"
			on:click={() => goto('/food-inventory')}
		>
			<IconX class="text-slate-600" />
		</button>
		<button type="submit" class="relative h-full aspect-square rounded-r flex justify-center items-center search-icon">
			<IconSearch class="text-slate-600" />
		</button>
	</form>
	<div class="relative w-full grid grid-cols-2 gap-5 overflow-auto pb-[90px]">
		{#each filteredFoods as food}
			<button on:click={() => (selectedFood = food)} class="relative w-full aspect-square rounded border border-r-red-200 p-5">
				<h2>{food.name}</h2>
				<span>{food.quantity} {food.unit} left in stock</span>
			</button>
		{:else}
			<p>No food added yet...</p>
		{/each}
	</div>
	<button
		type="button"
		on:click={() => (show = true)}
		class="fixed bottom-5 right-5 w-[50px] aspect-square rounded flex justify-center items-center bg-red-400 text-white"
	>
		<IconPlus />
	</button>
</main>
<AddModal bind:show />
<EditModal food={selectedFood} />

<style lang="postcss">
	.search-icon::before {
		content: '';
		position: absolute;
		top: 25%;
		left: 0;
		width: 1px;
		height: 50%;
		background-color: rgb(148, 163, 184);
	}
</style>
