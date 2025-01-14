// Just for easier testing. This is a lib mainly
package main

import (
    "os"
    "time"

    c "github.com/helto4real/go-hassclient/client"
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log *logrus.Entry

func main() {

    osSignal := make(chan os.Signal, 1)

    cl := c.NewHassClient()

    go cl.Start("localhost:8124", false, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIwNWYyOGFlZjkzYWU0MDUwODljMTYyYWNmNzZmNDk2MCIsImlhdCI6MTcwOTc4OTQ2OSwiZXhwIjoyMDI1MTQ5NDY5fQ.BqUndJESpV8jLTZO8aAT2uqDxmCb5cR0ZhJDewHQPAY")
    timeTickCheckState := time.Tick(5 * time.Second)
    // timeTickToggle := time.Tick(10 * time.Second)
    for {
        select {
        // case ev, mc := <-cl.HassCallServiceEventChannel:
        //  //log.Printf("%v", message)

        case _, mc := <-cl.GetHassChannel():
            if !mc {
                log.Println("Main channel terminating, exiting Loop")
                return
            }
        //  switch m := message.(type) {
        //  case c.HassEntity:
        //      //log.Info(m)
        //  case c.HassCallServiceEvent:
        //      if m.Service == "turn_on" {

        //          entityID, _ := m.ServiceData["entity_id"]

        //          if entityID == "switch.switcher" {
        //              cl.SetEntity(&c.HassEntity{ID: "switch.switcher", Name: "switch.switcher",
        //                  New: c.HassEntityState{State: "on",
        //                      Attributes: map[string]interface{}{"battery_level": "75", "icon": "mdi:pi-box"}}})
        //          }

        //      } else if m.Service == "turn_off" {

        //          entityID, _ := m.ServiceData["entity_id"]

        //          if entityID == "switch.switcher" {
        //              cl.SetEntity(&c.HassEntity{ID: "switch.switcher", Name: "switch.switcher",
        //                  New: c.HassEntityState{State: "off",
        //                      Attributes: map[string]interface{}{"battery_level": "75", "icon": "mdi:pi-box"}}})
        //          }

        //      }
        //  default:
        //      log.Warn(m)
        //  }

        case <-osSignal:
            log.Println("OS SIGNAL")
            cl.Stop()

        // case <-timeTickToggle:
        //  log.Print("Stopping...")
        //  cl.Stop()
        //  return
        case <-timeTickCheckState:
            //entity, ok := cl.GetEntity("light.tomas_rum_fonster")
            //if ok {

            //log.Printf("Light has the state: %s", entity.New.State)
            // cl.SetEntity(&c.HassEntity{ID: "switch.switcher", Name: "switch.switcher", New: c.HassEntityState{
            //  State:      "off",
            //  Attributes: map[string]interface{}{"battery_level": "75", "icon": "mdi:pi-box"}}})
            //return
            //cl.CallService("turn_on", map[string]string{"entity_id": "switch.switcher"})
            //return
            //}
        }

    }
}

func init() {
    log = logrus.WithField("prefix", "hassclient")
    Formatter := new(prefixed.TextFormatter)
    Formatter.FullTimestamp = true
    Formatter.TimestampFormat = "2006-01-02 15:04:05"
    Formatter.ForceColors = false
    Formatter.ForceFormatting = false
    logrus.SetFormatter(Formatter)
    logrus.SetLevel(logrus.TraceLevel)
    //  log.Level = logrus.DebugLevel
}
