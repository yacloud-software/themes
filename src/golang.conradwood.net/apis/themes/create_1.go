// client create: ThemesClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/themes/themes.proto
   gopackage : golang.conradwood.net/apis/themes
   importname: ai_0
   clientfunc: GetThemes
   serverfunc: NewThemes
   lookupfunc: ThemesLookupID
   varname   : client_ThemesClient_0
   clientname: ThemesClient
   servername: ThemesServer
   gsvcname  : themes.Themes
   lockname  : lock_ThemesClient_0
   activename: active_ThemesClient_0
*/

package themes

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_ThemesClient_0 sync.Mutex
  client_ThemesClient_0 ThemesClient
)

func GetThemesClient() ThemesClient { 
    if client_ThemesClient_0 != nil {
        return client_ThemesClient_0
    }

    lock_ThemesClient_0.Lock() 
    if client_ThemesClient_0 != nil {
       lock_ThemesClient_0.Unlock()
       return client_ThemesClient_0
    }

    client_ThemesClient_0 = NewThemesClient(client.Connect(ThemesLookupID()))
    lock_ThemesClient_0.Unlock()
    return client_ThemesClient_0
}

func ThemesLookupID() string { return "themes.Themes" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("themes.Themes")
   AddService("themes.Themes")
}
