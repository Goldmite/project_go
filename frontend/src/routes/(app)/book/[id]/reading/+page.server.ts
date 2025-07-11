import type { Actions, PageServerLoad } from "./$types";


export const load: PageServerLoad = async ({ params }) => {
    
    return ;
}

export const actions = {
    default: async(event) => {
        
        return { success: true };
    }
} satisfies Actions;