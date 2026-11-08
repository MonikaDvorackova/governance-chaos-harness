package main
import("fmt";"math/rand";"time")
type Event struct{Gate string; Approved bool}
func allowed(events []Event)bool{g:=map[string]bool{};for _,e:=range events{if !e.Approved{return false};g[e.Gate]=true};return g["evaluation"]&&g["risk"]&&g["human"]}
func main(){r:=rand.New(rand.NewSource(17));gates:=[]string{"evaluation","risk","human"};violations:=0;for trial:=0;trial<50000;trial++{n:=r.Intn(8);xs:=make([]Event,n);for i:=range xs{xs[i]=Event{gates[r.Intn(3)],r.Intn(5)!=0}};if allowed(xs){seen:=map[string]bool{};for _,x:=range xs{seen[x.Gate]=seen[x.Gate]||x.Approved};if !seen["evaluation"]||!seen["risk"]||!seen["human"]{violations++}}};fmt.Printf("trials=50000 invariant_violations=%d seed=17\n",violations);if violations>0{panic("gate invariant violated")};_ = time.Now()}
