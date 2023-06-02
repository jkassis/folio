interface Query {
  /* You dont need to worry about the specifics here */
}

interface SearchResult {
  /* You dont need to worry about the specifics here */
}

interface Page {
  results: SearchResult[]
  hasMoreResults: boolean
}

interface SearchService {
  /**
   * Contract:
   * Returns the next pageSize results starting from startFrom (non-inclusive).
   * May return fewer than pageSize results if there aren't enough remaining,
   * i.e. on the last page.
   *
   * If the returned page's hasMoreResults flag is set, it is guaranteed
   * that a subsequent search with the same query, starting from the last
   * result in the returned page, will contain at least one result.
   */
  search(query: Query, pageSize: number, startFrom?: SearchResult): Page
}

export function getAllResults(
  service: SearchService,
  query: Query
): SearchResult[] {
  const pageSize = 50 // let's fetch 50 results at a time

  // TODO fetch all results and return them
  let allResults: SearchResult[] = []
  var nextStart: SearchResult = {}
  while (true) {
    var page = service.search(query, pageSize, nextStart)
    allResults = allResults.concat(page.results)
    if (!page.hasMoreResults) break
    nextStart = page.results[page.results.length - 1]
  }

  return allResults
}

// TODO what's wrong with this FilteredSearchService?
type Filter = (result: SearchResult) => boolean
export class FilteredSearchService implements SearchService {
  private unfiltered: SearchService
  private filter: Filter

  search(query: Query, pageSize: number, startFrom?: SearchResult): Page {
    let page = this.unfiltered.search(query, pageSize, startFrom)

    let filteredResults: SearchResult[] = []
    page.results.forEach((result) => {
      if (this.filter(result)) {
        filteredResults.push(result)
      }
    })

    return {
      results: filteredResults,
      hasMoreResults: page.hasMoreResults,
    }
  }
}


// TODO fix the problems with the previous FilteredSearchService?
export class NewFilteredSearchService implements SearchService {
  private unfiltered: SearchService
  private filter: Filter

  search(query: Query, pageSize: number, startFrom?: SearchResult): Page {
    let filteredResults: SearchResult[] = []
    var nextStart: SearchResult | undefined = startFrom
    var added: number = 0
    let hasMoreResults: boolean | null = null
    while (hasMoreResults == null) {
      let unfilteredPage = this.unfiltered.search(query, pageSize, nextStart)
      if (unfilteredPage.results.length == 0) {
        hasMoreResults = false
        continue
      }
      nextStart = unfilteredPage[unfilteredPage.results.length - 1]
      unfilteredPage.results.forEach(result => {
        if (this.filter(result)) {
          if (added < pageSize) {
            added++
            filteredResults.push(result)
          } else {
            hasMoreResults = true
          }
        }
      })

      if (hasMoreResults == null && !unfilteredPage.hasMoreResults) {
        hasMoreResults = false
      }
    }

    return {
      results: filteredResults,
      hasMoreResults: hasMoreResults
    }
  }
}

