<script lang="ts">
	import { enhance } from '$app/forms';
	import { addToast } from '$stores/toast';
	import { fade, fly } from 'svelte/transition';
	import type { SubmitFunction } from './$types';
	import { foods } from '$stores/foods';

	export let show: boolean;

	const customEnhance: SubmitFunction = () => {
		return async ({ result }) => {
			if (result.type === 'failure' && result.data) {
				return addToast({ type: 'error', message: result.data.message });
			} else if (result.type === "success" && result.data?.method === "add") {
                addToast({ type: "success", message: result.data.message })
                $foods = [...$foods, result.data.data];
                show = false;
            }
		};
	};
</script>

{#if show}
	<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-static-element-interactions -->
	<div
		on:click|self={() => (show = false)}
		transition:fade
		class="fixed top-0 left-0 w-full h-full flex justify-center items-center bg-[rgba(0,0,0,0.1)] z-50"
	>
		<article
			class="relative w-full mx-5 bg-white rounded-lg p-5 shadow-lg"
			transition:fly={{ y: 50 }}
		>
			<form
				method="POST"
                action="?/add"
				class="relative w-full flex flex-col items-center gap-5"
				use:enhance={customEnhance}
			>
				<h2 class="text-xl">Add a new food</h2>
				<input
					name="name"
					placeholder="Food name..."
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
				<input
					name="quantity"
					placeholder="Food quantity"
					type="number"
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
                <input
					name="unit"
					placeholder="Food unit"
					class="relative w-full h-[50px] rounded bg-slate-200 px-5"
				/>
				<button type="submit" class="relative w-full h-[50px] rounded bg-red-400 text-white">
					Add food
				</button>
			</form>
		</article>
	</div>
{/if}
