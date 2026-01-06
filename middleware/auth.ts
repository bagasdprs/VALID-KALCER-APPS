export default defineNuxtRouteMiddleware((to, from) => {
  // 1. Cek apakah ada cookie token (Tanda pengenal user)
  const token = useCookie("token");

  // 2. Kalau token gak ada (artinya belum login)
  if (!token.value) {
    // Redirect paksa ke halaman login
    return navigateTo("/login");
  }

  // Kalau ada token, silakan lanjut (Pass)
});
