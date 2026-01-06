export default defineNuxtRouteMiddleware((to, from) => {
  // 1. Ambil data Token dan Role dari Cookie
  const token = useCookie("token");
  const role = useCookie("role"); // Kita simpan role di cookie juga biar gampang

  // 2. Cek Login dulu: Kalau gak ada token, suruh login
  if (!token.value) {
    return navigateTo("/login");
  }

  // 3. Cek Role: Kalau login tapi BUKAN admin
  if (role.value !== "admin") {
    // Tendang balik ke Home (atau halaman 403 Forbidden)
    return navigateTo("/");
  }

  // Kalau token ada DAN role-nya admin, silakan masuk Yang Mulia 👑
});
