package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	// Essa é a 1ª vez que usamos o mutex.
	// Ele, essencialmente, trava o servidor para que possamos
	// manipular mapa dos posts sem nos preocupar com outra
	// requisição tentando fazer a mesma coisa ao mesmo tempo
	postMu.Lock()
	defer postMu.Unlock()

	ps := make([]*Post, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, &p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ps)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro ao ler o body da requisição", http.StatusBadRequest)
		return
	}	

	var p Post
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "erro ao ler o body da requisição", http.StatusBadRequest)
		return
	}

	postMu.Lock()
	defer postMu.Unlock()

	p.ID = nextID
	nextID++
	posts[p.ID] = p

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&p)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postMu.Lock()
	defer postMu.Unlock()

	p, ok := posts[id]
	if !ok {
		http.Error(w, "post não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&p)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postMu.Lock()
	defer postMu.Unlock()

	_, ok := posts[id]
	if !ok {
		http.Error(w, "post não encontrado", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusOK)
}
