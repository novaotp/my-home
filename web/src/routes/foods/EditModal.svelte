<script lang="ts">
	import { enhance } from '$app/forms';
	import { addToast } from '$stores/toast';
	import { fade, fly } from 'svelte/transition';
	import type { SubmitFunction } from './$types';
	import { foods } from '$stores/foods';
	import type { Food } from '$models/Food';

	export let food: Food | null;

	const editEnhance: SubmitFunction = () => {
		return async ({ result }) => {
			if (result.type === 'failure' && result.data) {
				return addToast({ type: 'error', message: result.data.message });
			} else if (result.type === "success" && result.data?.method === "edit") {
				const editedFood = result.data.data;
                addToast({ type: "success", message: result.data.message })
                $foods = $foods.map(f => f.id === editedFood.id ? editedFood : f);
                food = null;
            }
		};
	};

	const deleteEnhance: SubmitFunction = () => {
		return async ({ result }) => {
			if (result.type === 'failure' && result.data) {
				return addToast({ type: 'error', message: result.data.message });
			} else if (result.type === "success" && result.data?.method === "delete") {
				const deletedFood = result.data.data;
                addToast({ type: "success", message: result.data.message })
                $foods = $foods.filter(f => f.id !== deletedFood.id);
                food = null;
            }
		};
	};
</script>

{#if food}
	<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
	<div
		on:click|self={() => (food = null)}
		transition:fade
		class="fixed top-0 left-0 w-full h-full flex justify-center items-center bg-[rgba(0,0,0,0.1)] z-50"
	>
		<article
			class="relative w-full mx-5 bg-white rounded-lg p-5 shadow-lg"
			transition:fly={{ y: 50 }}
		>
			<form
				method="POST"
                action="?/edit"
				class="relative w-full flex flex-col items-center gap-5"
				use:enhance={editEnhance}
			>
				<h2 class="text-xl">Edit a food</h2>
				<input type="hidden" name="id" value={food.id} />
				<input
					name="name"
					value={food.name}
					placeholder="Food name..."
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
				<input
					name="quantity"
					value={food.quantity}
					placeholder="Food quantity"
					type="number"
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
                <input
					name="unit"
					value={food.unit}
					placeholder="Food unit"
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
				<div class="relative w-full flex gap-5">
					<form method="POST" action="?/delete" class="relative w-full h-[50px]">
						<input type="hidden" name="id" value={food.id} />
						<button type="submit" class="relative w-full h-full rounded text-slate-500 border border-slate-500">
							Delete food
						</button>
					</form>
					<button type="submit" class="relative w-full h-[50px] rounded bg-red-400 text-white">
						Edit food
					</button>
				</div>
			</form>
		</article>
	</div>
{/if}
