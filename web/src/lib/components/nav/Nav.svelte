<script lang="ts">
	import MenuIcon from '@tabler/icons-svelte/IconMenu.svelte';
	import CloseIcon from '@tabler/icons-svelte/IconX.svelte';
	import IconLogout from '@tabler/icons-svelte/IconLogout.svelte';
	import { FullScreen, Card } from '$components/ui';
	import { links } from './Nav';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let showMenu: boolean = false;
</script>

<!-- Mobile nav with modal -->
<div class="md:hidden block text-neutral-950">
	<nav class="relative w-full h-[60px] flex justify-start px-5 py-[10px]">
		<button
			on:click={() => (showMenu = true)}
			class="relative h-full flex justify-center items-center"
		>
			<MenuIcon />
		</button>
	</nav>
	{#if showMenu}
		<FullScreen.Backdrop
			class="flex justify-center items-center"
			on:click={() => (showMenu = false)}
		>
			<Card class="mx-auto w-[calc(100%-40px)] flex flex-col gap-4" options={{ y: 50 }}>
				<header class="relative flex w-full items-center justify-between">
					<h2 class="text-xl">Menu</h2>
					<button
						on:click={() => (showMenu = false)}
						class="relative h-full flex justify-center items-center"
					>
						<CloseIcon />
					</button>
				</header>
				<menu>
					{#each links as { href, label, icon }}
						<li
							class="relative flex h-[50px] w-full items-center justify-center
                                    border-x border-t border-gray-400 first-of-type:rounded-t-md
                                    last-of-type:rounded-b-md last-of-type:border-b"
						>
							<button
								on:click={() => {
									if ($page.url.pathname !== href) {
										goto(href);
									}
									showMenu = false;
								}}
								class="relative flex h-full w-full items-center justify-start px-5 gap-5"
							>
								<svelte:component this={icon} />
								<span>{label}</span>
							</button>
						</li>
					{/each}
				</menu>
			</Card>
		</FullScreen.Backdrop>
	{/if}
</div>

<!-- Desktop nav -->
<nav class="hidden md:flex flex-col justify-between items-center h-full w-[100px] p-5 bg-stone-300">
	<ul class="relative w-full flex flex-col gap-3">
		{#each links as { href, label, icon }}
			<li class="relative flex w-full aspect-square items-center justify-center">
				<button
					on:click={() => {
						if ($page.url.pathname !== href) {
							goto(href);
						}
					}}
					class="relative flex flex-col h-full w-full items-center justify-evenly gap-1 p-2
                            text-center rounded-md hover:bg-stone-400 text-xs"
				>
					<svelte:component this={icon} />
					<span>{label}</span>
				</button>
			</li>
		{/each}
	</ul>
	<ul class="relative w-full flex flex-col gap-3">
		<li class="relative flex w-full aspect-square items-center justify-center">
			<a
				href="/auth/logout"
				class="relative flex flex-col h-full w-full items-center justify-evenly gap-1 p-2 text-center
                        rounded-md bg-red-500 duration-150 ease-linear hover:bg-red-600 text-xs"
			>
				<IconLogout class="text-white" />
			</a>
		</li>
	</ul>
</nav>
