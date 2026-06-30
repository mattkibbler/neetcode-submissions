
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var builder strings.Builder
    
    for i := 0; i < len(strs); i++ {
        builder.WriteString(strconv.Itoa(len(strs[i])))
        builder.WriteString("#")
        builder.WriteString(strs[i])
    }

    return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
    result := make([]string, 0)
    i := 0

    var lenStrBuilder strings.Builder
    for i < len(encoded) {
        if encoded[i] != '#' {
            lenStrBuilder.WriteByte(encoded[i])
            i++
        } else if encoded[i] == '#' { 
            i += 1 // increment to skip over #
            lenNumber, _ := strconv.Atoi(lenStrBuilder.String())
            result = append(result, encoded[i:i+lenNumber])
            i += lenNumber
            lenStrBuilder.Reset()
        }        
    }
    return result
}
