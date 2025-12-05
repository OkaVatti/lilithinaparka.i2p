import { marked } from 'marked'

export const useMarkdown = () => {
  const renderMarkdown = (content: string): string => {
    if (!content) return ''
    
    // Configure marked options
    marked.setOptions({
      breaks: true,
      gfm: true
    })
    
    return marked.parse(content) as string
  }
  
  return {
    renderMarkdown
  }
}