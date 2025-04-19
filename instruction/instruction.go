package instruction

// Instruction é o tipo que descreve a configuração da requisição
type Instruction struct {
    Endpoint    string            // URL do endpoint
    Method      string            // Método HTTP (GET, POST, PUT, etc.)
    Headers     map[string]string // Headers (ex.: autenticação, content-type)
    Fields      []string          // Campos a serem extraídos da resposta
    RequestBody []byte            // Corpo da requisição (para POST/PUT)
    AuthType    string            // Tipo de autenticação (Bearer, Basic, etc.)
    AuthValue   map[string]string // Autenticação (ex: username:calvin, password:senha, token:xxx)
}

// NewInstruction cria uma nova instância de Instruction com valores padrão
func NewInstruction(endpoint, method string, headers map[string]string, fields []string, requestBody []byte, authType string, authValue map[string]string) *Instruction {
    if method == "" {
        method = "GET" // Método padrão
    }

    return &Instruction{
        Endpoint:    endpoint,
        Method:      method,
        Headers:     headers,
        Fields:      fields,
        RequestBody: requestBody,
        AuthType:    authType,
        AuthValue:   authValue,
    }
}
