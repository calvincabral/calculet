package filter

import "strings"

// ApplyFilter aplica um filtro nos dados extraídos da resposta, usando um operador básico
func ApplyFilter(data map[string]interface{}, filter string) map[string]interface{} {
    filteredData := make(map[string]interface{})

    // Aplicação simples de filtro (exemplo: "field1=value")
    filterParts := strings.Split(filter, "=")
    if len(filterParts) == 2 {
        field := filterParts[0]
        value := filterParts[1]

        if val, exists := data[field]; exists && val == value {
            filteredData[field] = val
        }
    }

    return filteredData
}
