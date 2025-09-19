import { c as create_ssr_component } from "../../../chunks/ssr.js";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `${$$result.head += `<!-- HEAD_svelte-145rdro_START -->${$$result.title = `<title>Perfil - EcoLink</title>`, ""}<!-- HEAD_svelte-145rdro_END -->`, ""} ${`<div class="text-center py-8" data-svelte-h="svelte-a0rb2h"><div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-eco-600"></div> <p class="mt-2 text-gray-600">Carregando perfil...</p></div>`}`;
});
export {
  Page as default
};
