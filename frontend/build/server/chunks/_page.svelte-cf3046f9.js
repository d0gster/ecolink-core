import { c as create_ssr_component, a as subscribe, d as each, e as escape } from './ssr-98fcd782.js';
import { a as isAuthenticated } from './auth-aa1f0a46.js';
import './index-a2bcafe3.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$unsubscribe_isAuthenticated;
  $$unsubscribe_isAuthenticated = subscribe(isAuthenticated, (value) => value);
  const providers = [
    {
      name: "Google",
      icon: "🔍",
      color: "bg-red-500 hover:bg-red-600",
      functional: true
    },
    {
      name: "Microsoft",
      icon: "🪟",
      color: "bg-blue-500 hover:bg-blue-600",
      functional: false
    },
    {
      name: "LinkedIn",
      icon: "💼",
      color: "bg-blue-700 hover:bg-blue-800",
      functional: false
    },
    {
      name: "GitHub",
      icon: "🐙",
      color: "bg-gray-800 hover:bg-gray-900",
      functional: false
    }
  ];
  $$unsubscribe_isAuthenticated();
  return `${$$result.head += `<!-- HEAD_svelte-1ihkjo1_START -->${$$result.title = `<title>Login - EcoLink</title>`, ""}<!-- HEAD_svelte-1ihkjo1_END -->`, ""} <div class="min-h-[80vh] flex items-center justify-center"><div class="max-w-md w-full bg-white rounded-xl shadow-lg p-8"><div class="text-center mb-8" data-svelte-h="svelte-4b5yz0"><div class="text-4xl mb-4">🌱</div> <h1 class="text-2xl font-bold text-gray-900 mb-2">Enter EcoLink</h1> <p class="text-gray-600">Log in to access your sustainable links</p></div> <div class="space-y-3">${each(providers, (provider) => {
    return `<button ${""} class="${"w-full flex items-center justify-center space-x-3 py-3 px-4 rounded-lg text-white font-medium transition-colors disabled:opacity-50 " + escape(provider.color, true) + " " + escape(!provider.functional ? "opacity-75" : "", true)}"><span class="text-xl">${escape(provider.icon)}</span> <span>${escape(`Continue with ${provider.name}`)} ${escape(!provider.functional ? " (Soon)" : "")}</span> </button>`;
  })}</div> <div class="mt-8 text-center text-sm text-gray-500" data-svelte-h="svelte-9ckleo"><p>By continuing, you agree to our</p> <p><a href="/terms" class="text-eco-600 hover:text-eco-700">Terms of Use</a> and
				<a href="/privacy" class="text-eco-600 hover:text-eco-700">Privacy Policy</a></p></div></div></div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-cf3046f9.js.map
