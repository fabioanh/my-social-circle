<script lang="ts">
    import { page } from '$app/stores';
    import { api } from '$lib/api';
    import type { Person, Fact } from '$lib/types';
    import { goto } from '$app/navigation';

    let person: Person | null = null;
    let facts: Fact[] = [];
    let errorMessage = '';
    let newFactContent = '';
  
    $: personId = $page.params.id as string;

    let isEditing = false;
    let editableName = '';

    // Reactive load to handle ID changes
    $: if (personId) {
        loadData(personId);
    }

    async function loadData(id: string) {
        if (!id) return;
        try {
            person = await api.getPerson(id);
            if (person && person.facts) {
                // Determine if facts are embedded or need separate fetch
                // The current API design in `GetPerson` fetches facts and embeds them.
                facts = person.facts;
                editableName = person.name;
            } else {
                facts = [];
            }
        } catch (e) {
            errorMessage = 'Failed to load person details.';
        }
    }

    const updatePerson = async () => {
        if (!editableName.trim() || !personId) return;
        try {
            await api.updatePerson(personId, { name: editableName });
            if (person) {
                person.name = editableName;
            }
            isEditing = false;
        } catch (e) {
            console.error('Failed to update person:', e);
            errorMessage = 'Failed to update person.';
        }
    };

    const addFact = async () => {
        if (!newFactContent.trim() || !personId) return;
        try {
            const f = await api.addFact(personId, newFactContent);
            facts = [f, ...facts];
            newFactContent = '';
        } catch (e) {
            errorMessage = 'Failed to add fact.';
        }
    };

    const deleteFact = async (factId: string) => {
        if (!personId) return;
        try {
            await api.deleteFact(personId, factId);
            facts = facts.filter(f => f.id !== factId);
        } catch (e) {
            console.error('Failed to delete fact:', e);
            errorMessage = 'Failed to delete fact.';
        }
    };

    function goBack() {
        if (person?.group_id) {
            goto(`/groups/${person.group_id}`);
        } else {
            goto('/dashboard');
        }
    }
</script>

<div class="max-w-3xl mx-auto p-6 space-y-8">
    <button on:click={goBack} class="text-purple-400 hover:text-purple-300 flex items-center gap-2 mb-4 transition-colors">
        &larr; Back to Circle
    </button>
    
    {#if person}
        <header class="border-b border-slate-700 pb-6 mb-8">
            {#if isEditing}
                <div class="space-y-4">
                    <input 
                        bind:value={editableName}
                        class="text-4xl font-extrabold bg-slate-900 border border-purple-500 rounded-lg px-4 py-2 w-full text-white outline-none focus:ring-2 focus:ring-purple-500"
                    />
                    <div class="flex gap-2">
                        <button on:click={updatePerson} class="bg-purple-600 hover:bg-purple-500 text-white px-6 py-2 rounded-lg font-medium transition-colors">
                            Save
                        </button>
                        <button on:click={() => { isEditing = false; if (person) editableName = person.name; }} class="bg-slate-700 hover:bg-slate-600 text-white px-6 py-2 rounded-lg font-medium transition-colors">
                            Cancel
                        </button>
                    </div>
                </div>
            {:else}
                <div class="flex justify-between items-start">
                    <div>
                        <h1 class="text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600 mb-2">{person.name}</h1>
                        <p class="text-slate-500">Member since {new Date(person.created_at).toLocaleDateString()}</p>
                    </div>
                    <button 
                        on:click={() => isEditing = true}
                        class="text-slate-400 hover:text-purple-400 p-2 rounded-lg hover:bg-slate-800 transition-all flex items-center gap-2"
                    >
                        <span class="text-sm font-medium hidden md:inline">Edit Name</span>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                            <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                        </svg>
                    </button>
                </div>
            {/if}
        </header>

        <section class="space-y-6">
            <h2 class="text-2xl font-semibold text-slate-200">Facts & Memories</h2>
            
            <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50 backdrop-blur-sm shadow-lg">
                <textarea 
                    bind:value={newFactContent}
                    placeholder="Add a new fact (e.g., 'Loves spicy food', 'Met at the conference in 2023')..." 
                    class="w-full bg-slate-900 border border-slate-700 rounded-lg px-4 py-3 text-white focus:ring-2 focus:ring-purple-500 outline-none resize-none h-24 mb-3 transition-all focus:border-purple-500"
                    on:keydown={(e) => e.key === 'Enter' && !e.shiftKey && (e.preventDefault(), addFact())}
                ></textarea>
                <div class="flex justify-end">
                    <button 
                        on:click={addFact}
                        disabled={!newFactContent.trim()}
                        class="bg-purple-600 hover:bg-purple-500 disabled:opacity-50 disabled:cursor-not-allowed text-white px-6 py-2 rounded-lg font-medium transition-all shadow-[0_0_15px_rgba(147,51,234,0.3)] hover:shadow-[0_0_25px_rgba(147,51,234,0.5)] active:scale-95"
                    >
                        Remember Fact
                    </button>
                </div>
            </div>

            <div class="space-y-4">
                {#if facts.length === 0}
                    <div class="text-center py-12 text-slate-600 italic">
                        No facts recorded yet. Start building your map!
                    </div>
                {/if}
                {#each facts as fact}
                    <div class="group relative bg-gradient-to-br from-slate-800 to-slate-900 p-5 rounded-xl border border-slate-700 hover:border-purple-500/30 transition-all hover:translate-x-1 shadow-sm">
                        <div class="pr-8">
                            <p class="text-slate-200 text-lg leading-relaxed">{fact.content}</p>
                        </div>
                        <div class="absolute top-4 right-4 flex items-center gap-3">
                            <span class="text-xs text-slate-600 font-mono opacity-0 group-hover:opacity-100 transition-opacity hidden sm:inline">
                                {new Date(fact.created_at).toLocaleDateString()}
                            </span>
                            <button 
                                on:click={() => deleteFact(fact.id)}
                                class="text-slate-600 hover:text-red-500 opacity-0 group-hover:opacity-100 transition-opacity p-1"
                                title="Delete Fact"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                                    <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
                                </svg>
                            </button>
                        </div>
                    </div>
                {/each}
            </div>
        </section>

    {:else if errorMessage}
        <div class="p-4 bg-red-500/20 text-red-200 rounded-lg border border-red-500/50">
            {errorMessage}
        </div>
    {:else}
        <div class="flex items-center justify-center p-12">
            <div class="w-8 h-8 border-4 border-purple-500 border-t-transparent rounded-full animate-spin"></div>
        </div>
    {/if}
</div>
