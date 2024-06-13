import { API_URL } from "$env/static/private";
import type { APIResponse, WithData } from "$lib/models/Responses";
import type { Food } from "$lib/models/Food";
import type { Actions, PageServerLoad } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ url }) => {
    if (url.searchParams.get("search") === "") {
        redirect(303, "/food-inventory");
    }

    const response = await fetch(`${API_URL}/api/v1/foods`);
    const result: WithData<APIResponse, Food[]> = await response.json();
    
    return {
        foods: result.data
    }
};

export const actions: Actions = {
    add: async ({ request }) => {
        const form = await request.formData();
        const data = {
            name: form.get("name")?.toString(),
            quantity: form.get("quantity")?.toString(),
            unit: form.get("unit")?.toString()
        }

        if (data.name === "" || data.quantity === "") {
            return fail(422, { data, message: "Fill all the inputs." })
        }

        const response = await fetch(`${API_URL}/api/v1/foods`, {
            method: "POST",
            body: JSON.stringify(data),
            headers: {
                "accept": "application/json",
                "content-type": "application/json"
            }
        })
        const result: WithData<APIResponse, Food> = await response.json();
        return { method: "add", ...result } as const
    },
    edit: async ({ request }) => {
        const form = await request.formData();
        const data = {
            id: form.get("id")!.toString(),
            name: form.get("name")?.toString(),
            quantity: form.get("quantity")?.toString(),
            unit: form.get("unit")?.toString()
        }

        if (data.name === "" || data.quantity === "") {
            return fail(422, { data, message: "Fill all the inputs." })
        }

        const response = await fetch(`${API_URL}/api/v1/foods/${data.id}`, {
            method: "PUT",
            body: JSON.stringify(data),
            headers: {
                "accept": "application/json",
                "content-type": "application/json"
            }
        })
        const result: WithData<APIResponse, Food> = await response.json();
        return { method: "edit", ...result } as const
    },
    delete: async ({ request }) => {
        const form = await request.formData();
        const id = form.get("id")!.toString();

        const response = await fetch(`${API_URL}/api/v1/foods/${id}`, {
            method: "DELETE",
            headers: {
                "accept": "application/json"
            }
        })
        const result: WithData<APIResponse, Food> = await response.json();
        return { method: "delete", ...result } as const
    }
}
