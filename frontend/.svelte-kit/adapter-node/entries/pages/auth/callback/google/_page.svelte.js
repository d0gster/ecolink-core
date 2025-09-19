import { c as create_ssr_component, b as subscribe } from "../../../../../chunks/ssr.js";
import { p as page } from "../../../../../chunks/stores.js";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$unsubscribe_page;
  $$unsubscribe_page = subscribe(page, (value) => value);
  $$unsubscribe_page();
  return `${$$result.head += `<!-- HEAD_svelte-nnn3os_START -->${$$result.title = `<title>Processing Login - EcoLink</title>`, ""}<!-- HEAD_svelte-nnn3os_END -->`, ""} <div class="min-h-[80vh] flex items-center justify-center"><div class="text-center">${`<div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600 mb-4"></div> <h1 class="text-xl font-semibold text-gray-900 mb-2" data-svelte-h="svelte-1t9uj43">Processing login...</h1> <p class="text-gray-600" data-svelte-h="svelte-oxe331">Wait a moment</p>`}</div></div>`;
});
export {
  Page as default
};
