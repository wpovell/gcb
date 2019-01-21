#include <string.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <linux/wireless.h>
#include <netinet/in.h>
#include <unistd.h>
#include "ssid.h"

#include <stdio.h>

// Returns true if `intr` is up and running
int link_up(char *intr) {
    int sock = socket(AF_INET, SOCK_DGRAM, IPPROTO_IP);
    if (sock < 0)
      return -1;

    struct ifreq freq;
    memset(&freq, 0, sizeof(struct iwreq));
    strncpy(freq.ifr_name, intr, sizeof(freq.ifr_name));

    int rv = ioctl(sock, SIOCGIFFLAGS, &freq);
    close(sock);
    if (rv < 0) {
      return -1;
    }

    return (freq.ifr_flags & IFF_UP) && (freq.ifr_flags & IFF_RUNNING);
}

// Return string essid that `intr` is connected to
char* ssid(char* intr) {
  // Check if link is up first
  if (link_up(intr) <= 0) {
    return NULL;
  }

  int sock = socket(AF_INET, SOCK_DGRAM, 0);
  if (sock < 0) {
    return NULL;
  }

  char *essid = malloc(sizeof(char) * (IW_ESSID_MAX_SIZE+1));
  struct iwreq wreq;
  memset(&wreq, 0, sizeof(struct iwreq));
  strncpy(wreq.ifr_name, intr, sizeof(wreq.ifr_name));
  wreq.u.essid.pointer = essid;
  wreq.u.essid.length = IW_ESSID_MAX_SIZE;

  int rv = ioctl(sock ,SIOCGIWESSID, &wreq);
  close(sock);
  if (rv < 0) {
    return NULL;
  }

  essid[wreq.u.essid.length] = '\0';
  return essid;
}
