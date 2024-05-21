import type { ComponentType } from "svelte";
import IconBabyBottle from "@tabler/icons-svelte/IconBabyBottle.svelte"
import IconShoppingCart from "@tabler/icons-svelte/IconShoppingCart.svelte"
import IconHome from "@tabler/icons-svelte/IconHome.svelte"

interface Link {
    href: string,
    label: string,
    icon: ComponentType
}

export const links: Link[] = [
    {
        href: "/",
        label: "Home",
        icon: IconHome
    },
    {
        href: "/food-inventory",
        label: "Food Inventory",
        icon: IconBabyBottle
    },
    {
        href: "/shopping-lists",
        label: "Shopping List",
        icon: IconShoppingCart
    }
]
