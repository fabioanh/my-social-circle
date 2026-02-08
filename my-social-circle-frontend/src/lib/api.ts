import type { Group, Person, Fact } from './types';

const API_BASE = 'http://localhost:8081';

async function fetchJson<T>(url: string, options?: RequestInit): Promise<T> {
    try {
        const res = await fetch(`${API_BASE}${url}`, options);
        if (!res.ok) {
            console.error(`API Error on ${url}:`, res.status, res.statusText);
            const text = await res.text();
            console.error('Response body:', text);
            throw new Error(`API Error: ${res.statusText}`);
        }
        if (res.status === 204) {
            return null as unknown as T;
        }
        return res.json();
    } catch (e) {
        console.error('Network or Parse Error:', e);
        throw e;
    }
}

export const api = {
    // Groups
    getGroups: () => fetchJson<Group[]>('/groups'),
    getGroup: (id: string) => fetchJson<Group>(`/groups/${id}`),
    createGroup: (group: Partial<Group>) => fetchJson<Group>('/groups', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(group)
    }),
    updateGroup: (id: string, group: Partial<Group>) => fetchJson<void>(`/groups/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(group)
    }),

    // People
    getPeopleByGroup: (groupId: string) => fetchJson<Person[]>(`/groups/${groupId}/people`),
    getPerson: (id: string) => fetchJson<Person>(`/people/${id}`),
    createPerson: (person: Partial<Person>) => fetchJson<Person>('/people', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(person)
    }),
    updatePerson: (id: string, person: Partial<Person>) => fetchJson<void>(`/people/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(person)
    }),

    // Facts
    addFact: (personId: string, content: string) => fetchJson<Fact>(`/people/${personId}/facts`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content })
    }),
    deleteFact: (personId: string, factId: string) => fetchJson<void>(`/people/${personId}/facts/${factId}`, {
        method: 'DELETE'
    })
};
