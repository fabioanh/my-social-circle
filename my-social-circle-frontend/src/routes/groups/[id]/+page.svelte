<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { api } from '$lib/api';
    import type { Group, Person } from '$lib/types';
    import { goto } from '$app/navigation';

    let group: Group | null = null;
    let people: Person[] = [];
    let errorMessage = '';
    let newPersonName = '';
  
    $: groupId = $page.params.id as string;

    let isEditing = false;
    let editableName = '';
    let editableDescription = '';

    onMount(async () => {
        if (!groupId) return;
        try {
            group = await api.getGroup(groupId);
            people = await api.getPeopleByGroup(groupId);
            if (group) {
                editableName = group.name;
                editableDescription = group.description;
            }
        } catch (e) {
            errorMessage = 'Failed to load group details.';
        }
    });

    const updateGroup = async () => {
        if (!editableName.trim() || !groupId) return;
        try {
            await api.updateGroup(groupId, { name: editableName, description: editableDescription });
            if (group) {
                group.name = editableName;
                group.description = editableDescription;
            }
            isEditing = false;
        } catch (e) {
            console.error('Failed to update group:', e);
            errorMessage = 'Failed to update group.';
        }
    };

    const addPerson = async () => {
        if (!newPersonName.trim() || !groupId) return;
        try {
            const p = await api.createPerson({ 
                name: newPersonName, 
                group_id: groupId 
            });
            people = [p, ...people];
            newPersonName = '';
        } catch (e) {
            console.error('Failed to add person:', e);
            errorMessage = 'Failed to add person.';
        }
    };
</script>

<div class="p-6 max-w-4xl mx-auto space-y-8">
    <button on:click={() => goto('/dashboard')} class="text-cyan-400 hover:text-cyan-300 flex items-center gap-2 mb-4">
        &larr; Back to Dashboard
    </button>

    {#if group}
        <header class="border-b border-slate-700 pb-6 mb-8">
            {#if isEditing}
                <div class="space-y-4">
                    <input 
                        bind:value={editableName}
                        class="text-4xl font-extrabold bg-slate-900 border border-cyan-500 rounded-lg px-4 py-2 w-full text-white outline-none focus:ring-2 focus:ring-cyan-500"
                    />
                    <textarea 
                        bind:value={editableDescription}
                        class="text-slate-400 text-lg bg-slate-900 border border-slate-700 rounded-lg px-4 py-2 w-full outline-none focus:ring-2 focus:ring-cyan-500 h-24"
                    ></textarea>
                    <div class="flex gap-2">
                        <button on:click={updateGroup} class="bg-cyan-600 hover:bg-cyan-500 text-white px-6 py-2 rounded-lg font-medium transition-colors">
                            Save
                        </button>
                        <button on:click={() => { if (group) { isEditing = false; editableName = group.name; editableDescription = group.description; } }} class="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg font-medium transition-colors">
                            Cancel
                        </button>
                    </div>
                </div>
            {:else}
                <div class="flex justify-between items-start">
                    <div>
                        <h1 class="text-4xl font-extrabold text-white mb-2">{group.name}</h1>
                        <p class="text-slate-400 text-lg">{group.description}</p>
                    </div>
                    <button 
                        on:click={() => isEditing = true}
                        class="text-slate-400 hover:text-cyan-400 p-2 rounded-lg hover:bg-slate-800 transition-all flex items-center gap-2"
                    >
                        <span class="text-sm font-medium hidden md:inline">Edit Circle</span>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                        </svg>
                    </button>
                </div>
            {/if}
        </header>

        <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50 backdrop-blur-sm">
            <h2 class="text-xl font-semibold mb-4 text-slate-200">Add Person to Circle</h2>
        <div class="flex flex-col md:flex-row gap-4">
                <input 
                    bind:value={newPersonName}
                    placeholder="Person Name (e.g., Alice Smith)" 
                    class="w-full md:flex-1 bg-slate-900 border border-slate-700 rounded-lg px-4 py-2 text-white focus:ring-2 focus:ring-cyan-500 outline-none"
                    on:keydown={(e) => e.key === 'Enter' && addPerson()}
                />
                <button 
                    on:click={addPerson}
                    class="w-full md:w-auto bg-cyan-600 hover:bg-cyan-500 text-white px-6 py-2 rounded-lg font-medium transition-colors"
                >
                    Add
                </button>
            </div>
        </div>

        <div class="space-y-4">
            <h2 class="text-2xl font-bold text-white mb-4">People in this Circle</h2>
            {#if people.length === 0}
                <p class="text-slate-500">No people added yet.</p>
            {/if}
            {#each people as person}
                <a href="/people/{person.id}" class="block group">
                    <div class="bg-slate-800 p-4 rounded-xl border border-slate-700 hover:border-purple-500/50 transition-all hover:bg-slate-750 flex justify-between items-center group-hover:translate-x-1">
                        <div>
                            <h3 class="text-xl font-bold text-white group-hover:text-purple-400 transition-colors">{person.name}</h3>
                            <p class="text-slate-500 text-sm">Since {new Date(person.created_at).toLocaleDateString()}</p>
                            {#if person.first_fact}
                                <div class="mt-2 text-slate-300 italic text-sm border-l-2 border-purple-500/30 pl-3">
                                    "{person.first_fact.content}"
                                </div>
                            {/if}
                        </div>
                        <div class="text-slate-600 group-hover:text-purple-400 transition-colors">
                            View details &rarr;
                        </div>
                    </div>
                </a>
            {/each}
        </div>

    {:else if errorMessage}
        <div class="p-4 bg-red-500/20 text-red-200 rounded-lg">
            {errorMessage}
        </div>
    {:else}
        <div class="text-slate-500 animate-pulse">Loading group details...</div>
    {/if}
</div>
