<script lang="ts">
	import IconPlus from '@tabler/icons-svelte/IconPlus.svelte';
	import IconSearch from '@tabler/icons-svelte/IconSearch.svelte';
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

	$: if ($page.url.searchParams.get("search") === "") {
		goto("/foods")
	}

	$: search = $page.url.searchParams.get("search") ?? "";
	$: filteredFoods = $foods.filter(food => search === "" || food.name.toLowerCase().includes(search.toLowerCase()))
</script>

<main class="relative w-full h-full flex flex-col p-5 gap-5">
	<h1>Foods here</h1>
	<form action="/foods" class="relative w-full min-h-[50px] flex">
		<input type="search" name="search" placeholder="Search food" class="relative w-full h-full px-5 rounded-l bg-slate-200" />
		<button type="submit" class="relative h-full aspect-square rounded-r bg-slate-200 flex justify-center items-center search-icon">
			<IconSearch class="text-slate-600" />
		</button>
	</form>
	<div class="relative w-full grid grid-cols-2 gap-5 overflow-auto pb-[90px]">
		{#each filteredFoods as food}
			<button
				on:click={() => (selectedFood = food)}
				class="relative w-full aspect-square rounded border border-r-red-200 p-5"
			>
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
		content: "";
		position: absolute;
		top: 25%;
		left: 0;
		width: 1px;
		height: 50%;
		background-color: rgb(148, 163, 184);
	}
</style>
