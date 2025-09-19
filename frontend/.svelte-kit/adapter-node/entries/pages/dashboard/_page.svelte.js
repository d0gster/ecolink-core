import { c as create_ssr_component } from "../../../chunks/ssr.js";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `${$$result.head += `<!-- HEAD_svelte-1dmdila_START -->${$$result.title = `<title>Dashboard - EcoLink</title>`, ""}<!-- HEAD_svelte-1dmdila_END -->`, ""} <div><div class="mb-8" data-svelte-h="svelte-zbx3qx"><h1 class="text-3xl font-bold text-gray-900">Dashboard</h1> <p class="text-gray-700 mt-2">Manage your shortened links</p></div> ${`<div class="text-center py-8" data-svelte-h="svelte-3he77g"><div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-eco-600"></div> <p class="mt-2 text-gray-600">Loading your links...</p></div>`}</div>`;
});
export {
  Page as default
};
