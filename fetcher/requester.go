package requester

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
)

// DoRequest faz uma requisição HTTP e retorna o conteúdo JSON.
func DoRequest(url string, method string, headers map[string]string, body []byte) (map[string]interface{}, error) {
    // Cria a requisição com o método desejado
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        log.Printf("Erro ao criar requisição: %s", err)
        return nil, err
    }

    // Adiciona os headers à requisição
    for key, value := range headers {
        req.Header.Add(key, value)
    }

    // Se houver corpo (para POST, PUT, etc.), adiciona
    if body != nil {
        req.Body = ioutil.NopCloser(bytes.NewReader(body))
    }

    // Faz a requisição HTTP
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Falha ao fazer requisição: %s", err)
        return nil, err
    }
    defer resp.Body.Close()

    // Verifica o status da resposta
    if resp.StatusCode != 200 {
        log.Printf("Erro: Status %d ao acessar o endpoint", resp.StatusCode)
        return nil, fmt.Errorf("Erro: Status %d", resp.StatusCode)
    }

    // Lê o corpo da resposta
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Erro ao ler corpo da resposta: %s", err)
        return nil, err
    }

    // Deserializa o corpo em JSON
    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Printf("Erro ao deserializar JSON: %s", err)
        return nil, err
    }

    return result, nil
}