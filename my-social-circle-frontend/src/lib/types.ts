export interface Group {
    id: string;
    name: string;
    description: string;
    created_at: string;
}

export interface Person {
    id: string;
    name: string;
    group_id: string;
    created_at: string;
    facts: Fact[];
    first_fact?: Fact;
}

export interface Fact {
    id: string;
    content: string;
    person_id: string;
    created_at: string;
}
