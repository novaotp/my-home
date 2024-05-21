<script lang="ts">
    export let variant: "primary" | "secondary" | "tertiary" = "primary";
    export let size: "small" | "medium" = "medium";

    const sizeClasses = {
        "small": "px-3 py-1 text-sm",
        "medium": "px-5 py-2 text-md"
    };

    const colorClasses = {
        "primary": "text-gray-50 bg-green-500",
        "secondary": "text-red-500 bg-red-50 border border-red-500",
        "tertiary": "text-gray-500"
    }
</script>

<button class="relative rounded {sizeClasses[size]} {colorClasses[variant]} {$$restProps['class'] || ''}" on:click>
    <slot />
</button>
