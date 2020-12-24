package boinga

import (
    "testing"
    "fmt"
)

func TestConvertingToBoingaWorks(t *testing.T) {
    value := "this is another string that, hopefully, will show everything works."
    expected := "boInGA boiNgA BoiNgA BOinGA boinga BoiNgA BOinGA boinga BoingA bOINgA BOINgA boInGA boiNgA BoIngA bOinGA boinga BOinGA boInGA bOinGA BoiNgA bOINgA BOIngA boinga boInGA boiNgA BoingA boInGA boINga boinga boiNgA BOINgA boinGA BoIngA bOIngA BoInGA boINgA boINgA BoiNGA boINga boinga BOInGA BoiNgA boINgA boINgA boinga BOinGA boiNgA BOINgA BOInGA boinga BoIngA bOInGA BoIngA bOinGA BoiNGA boInGA boiNgA BoiNgA bOINgA BOIngA boinga BOInGA BOINgA bOinGA BOiNgA BOinGA bOINga"
    result := ToBoinga(value)
    if expected != result {
        t.Error(fmt.Sprintf("\n# Expected\n%s\n\n# Gotten\n%s\n\n", expected, result))
    }
}

func TestConvertingFromBoingaWorks(t *testing.T) {
    value := "boInGA boiNgA BoiNgA BOinGA boinga BoiNgA BOinGA boinga BoingA bOINgA BOINgA boInGA boiNgA BoIngA bOinGA boinga BOinGA boInGA bOinGA BoiNgA bOINgA BOIngA boinga boInGA boiNgA BoingA boInGA boINga boinga boiNgA BOINgA boinGA BoIngA bOIngA BoInGA boINgA boINgA BoiNGA boINga boinga BOInGA BoiNgA boINgA boINgA boinga BOinGA boiNgA BOINgA BOInGA boinga BoIngA bOInGA BoIngA bOinGA BoiNGA boInGA boiNgA BoiNgA bOINgA BOIngA boinga BOInGA BOINgA bOinGA BOiNgA BOinGA bOINga"
    expected := "THIS IS ANOTHER STRING THAT, HOPEFULLY, WILL SHOW EVERYTHING WORKS."
    result := FromBoinga(value)
    if expected != result {
        t.Error(fmt.Sprintf("\n# Expected\n%s\n\n# Gotten\n%s\n\n", expected, result))
    }
}
