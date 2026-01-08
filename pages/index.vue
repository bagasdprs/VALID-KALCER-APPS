<script setup lang="ts">
import ScannerHero from "~/components/skena/ScannerHero.vue";
import About from "~/components/landing/About.vue";
import Services from "~/components/landing/Services.vue";
import Pricing from "~/components/landing/Pricing.vue";
import Testimonials from "~/components/landing/Testimonials.vue";
import FAQ from "~/components/landing/FAQ.vue";

definePageMeta({
  layout: "default",
});
</script>

<template>
  <div>
    <ScannerHero />
    <About />
    <Services />
    <Pricing />
    <Testimonials />
    <FAQ />
  </div>
</template>
