from ..models.models import SearchData, SearchResponse, Source


def process(data: SearchData | None) -> SearchResponse | None:
    if data == None:
        return None

    response = SearchResponse(
        result="Fake", score=78, sources=[Source(url="http://example_source.com")]
    )

    return response
