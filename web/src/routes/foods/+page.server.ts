import { API_URL } from "$env/static/private";
import type { APIResponse, WithData } from "$lib/models/Responses";
import type { Food } from "$lib/models/Food";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
    const response = await fetch(`${API_URL}/api/v1/foods`);
    const result: WithData<APIResponse, Food[]> = await response.json();
    
    return {
        foods: result.data
    }
};
