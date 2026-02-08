<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { Group } from '$lib/types';

    let groups: Group[] = [];
    let errorMessage = '';

    onMount(async () => {
        try {
            groups = await api.getGroups();
        } catch (e) {
            errorMessage = 'Failed to load groups. Ensure backend is running.';
        }
    });

    // Simple form state
    let newGroupName = '';
    let newGroupDesc = '';

    const createGroup = async () => {
        try {
            const g = await api.createGroup({ name: newGroupName, description: newGroupDesc });
            groups = [...groups, g];
            newGroupName = '';
            newGroupDesc = '';
        } catch (e) {
            console.error('Failed to create group:', e);
            errorMessage = 'Failed to create group. Check console for details.';
        }
    };
</script>

<div class="p-6 max-w-6xl mx-auto space-y-8">
    <header class="flex justify-between items-center">
        <h1 class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-cyan-400 to-purple-600">Your Circles</h1>
        <!-- Add Group Modal Trigger could go here, keeping it inline for now -->
    </header>

    {#if errorMessage}
        <div class="p-4 bg-red-500/20 text-red-200 rounded-lg border border-red-500/50">
            {errorMessage}
        </div>
    {/if}

    <!-- Create Group Section -->
    <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50 backdrop-blur-sm">
        <h2 class="text-xl font-semibold mb-4 text-slate-200">Create New Circle</h2>
        <div class="flex flex-col md:flex-row gap-4">
            <input 
                bind:value={newGroupName}
                placeholder="Group Name (e.g., Hiking Club)" 
                class="w-full md:flex-1 bg-slate-900 border border-slate-700 rounded-lg px-4 py-2 text-white focus:ring-2 focus:ring-cyan-500 outline-none"
            />
            <input 
                bind:value={newGroupDesc}
                placeholder="Description" 
                class="w-full md:flex-[2] bg-slate-900 border border-slate-700 rounded-lg px-4 py-2 text-white focus:ring-2 focus:ring-cyan-500 outline-none"
            />
            <button 
                on:click={createGroup}
                class="w-full md:w-auto bg-cyan-600 hover:bg-cyan-500 text-white px-6 py-2 rounded-lg font-medium transition-colors"
            >
                Create
            </button>
        </div>
    </div>

    <!-- Groups Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {#each groups as group}
            <a href="/groups/{group.id}" class="block group">
                <div class="h-full bg-slate-800 p-6 rounded-2xl border border-slate-700 hover:border-cyan-500/50 transition-all hover:shadow-[0_0_20px_rgba(6,182,212,0.15)] hover:-translate-y-1">
                    <h3 class="text-2xl font-bold text-white mb-2 group-hover:text-cyan-400 transition-colors">{group.name}</h3>
                    <p class="text-slate-400">{group.description}</p>
                    <div class="mt-4 text-sm text-slate-500 flex items-center gap-2">
                        <span>Created {new Date(group.created_at).toLocaleDateString()}</span>
                    </div>
                </div>
            </a>
        {/each}
        
        {#if groups.length === 0 && !errorMessage}
            <div class="col-span-full text-center py-12 text-slate-500">
                No circles found. Create one to get started!
            </div>
        {/if}
    </div>
</div>
