<script>
	import ToastContainer from '$components/toast/ToastContainer.svelte';
	import Navigation from '$components/nav/Nav.svelte';
	import '../app.css';
</script>

<ToastContainer />
<div class="relative h-full w-full flex flex-col md:flex-row">
    <Navigation />
    <div class="relative w-full h-[calc(100%-60px)] md:h-full flex flex-col">
        <slot />
    </div>
</div>
