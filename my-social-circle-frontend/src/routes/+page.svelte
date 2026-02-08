<script lang="ts">
	// Simple landing page logic
</script>

<div class="min-h-screen bg-slate-900 text-white flex flex-col items-center justify-center p-4 selection:bg-cyan-500 selection:text-white">
	<div class="absolute inset-0 overflow-hidden pointer-events-none">
		<div class="absolute -top-[20%] -left-[10%] w-[50%] h-[50%] rounded-full bg-cyan-500/20 blur-[120px]"></div>
		<div class="absolute top-[40%] -right-[10%] w-[40%] h-[40%] rounded-full bg-purple-500/20 blur-[120px]"></div>
	</div>

	<main class="relative z-10 max-w-4xl w-full text-center space-y-8">
		<h1 class="text-5xl md:text-7xl font-extrabold tracking-tight bg-clip-text text-transparent bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 animate-gradient-x">
			My Social Circle
		</h1>
		<p class="text-xl md:text-2xl text-slate-300 max-w-2xl mx-auto leading-relaxed">
			Remember every connection. Identify faces, recall facts, and strengthen your relationships before every gathering.
		</p>
		
		<div class="flex flex-col sm:flex-row items-center justify-center gap-4 mt-8">
			<a href="/dashboard" class="px-8 py-4 bg-cyan-500 hover:bg-cyan-400 text-slate-900 font-bold rounded-full transition-all hover:scale-105 hover:shadow-[0_0_20px_rgba(6,182,212,0.5)] active:scale-95 text-center">
				Get Started
			</a>
			<a href="https://github.com/fabioanh/my-social-circle-backend" target="_blank" class="px-8 py-4 bg-slate-800 hover:bg-slate-700 text-cyan-400 border border-slate-700 hover:border-cyan-500/50 rounded-full font-semibold transition-all hover:scale-105 active:scale-95 text-center">
				Learn More
			</a>
		</div>
	</main>

	<footer class="absolute bottom-8 text-slate-500 text-sm">
		&copy; {new Date().getFullYear()} My Social Circle. All rights reserved.
	</footer>
</div>

<style>
	:global(body) {
		background-color: #0f172a; /* slate-900 */
		margin: 0;
		font-family: 'Inter', sans-serif;
	}
	
	.animate-gradient-x {
		background-size: 200% 200%;
		animation: gradient-x 6s ease infinite;
	}

	@keyframes gradient-x {
		0%, 100% {
			background-position: 0% 50%;
		}
		50% {
			background-position: 100% 50%;
		}
	}
</style>
