package handler

/* From 2013/08/28 - 2017/08/08, Maxmind provide GeoLite dataset in legacy format

gs://downloader-mlab-oti/Maxmind/2013/08/28/20130828T184800Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2013/09/07/20130907T170000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2013/10/07/20131007T170000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2013/11/07/20131107T160000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2013/12/07/20131207T160000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/02/07/20140207T160000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/03/07/20140307T160000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/04/07/20140407T170000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/05/04/20140504T070800Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/05/08/20140508T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/06/08/20140608T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/07/08/20140708T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/08/08/20140808T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/09/08/20140908T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/10/28/20141028T032100Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/11/08/20141108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2014/12/08/20141208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/01/08/20150108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/02/08/20150208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/03/08/20150308T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/04/08/20150408T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/05/08/20150508T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/06/08/20150608T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/07/08/20150708T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/08/08/20150808T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/09/08/20150908T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/10/08/20151008T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/11/03/20151103T204100Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/11/08/20151108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2015/12/08/20151208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/01/08/20160108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/02/08/20160208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/03/08/20160308T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/04/08/20160408T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/05/08/20160508T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/06/08/20160608T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/07/08/20160708T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/08/08/20160808T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/09/08/20160908T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/10/08/20161008T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/11/08/20161108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2016/12/08/20161208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/01/08/20170108T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/02/08/20170208T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/03/08/20170308T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/04/08/20170408T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/05/08/20170508T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/06/08/20170608T080000Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/07/05/20170705T153500Z-GeoLiteCity.dat.gz
gs://downloader-mlab-oti/Maxmind/2017/08/08/20170808T080000Z-GeoLiteCity.dat.gz

   Each data set cover the time range between the last available dataset and its own time stamp.
   There are IP v6 datasets as well.
 
   From 2017/08/15 - present, Maxmind provides both legacy format and GeoLite2

gs://downloader-mlab-oti/Maxmind/2017/08/15/20170815T200728Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/08/15/20170815T200946Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/09/01/20170901T004438Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/09/01/20170901T085053Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/09/07/20170907T023620Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/09/07/20170907T030659Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/10/01/20171001T003046Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/10/01/20171001T025802Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/10/04/20171004T191825Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/10/05/20171005T033334Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/10/05/20171005T040958Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/11/01/20171101T013013Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/11/06/20171106T232541Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/12/01/20171201T074609Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/12/01/20171201T183227Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2017/12/06/20171206T205411Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2018/01/01/20180101T033908Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2018/01/05/20180105T203044Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2018/02/01/20180201T045126Z-GeoLite2-City-CSV.zip
gs://downloader-mlab-oti/Maxmind/2018/02/08/20180208T013555Z-GeoLite2-City-CSV.zip
...


*/
import (
	"context"
	"log"
	"os"
	"regexp"

	"github.com/m-lab/annotation-service/loader"
	"github.com/m-lab/annotation-service/parser"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/storage"
)

// This is the regex used to filter for which files we want to consider acceptable for using with Geolite2
var GeoLegacyRegex = regexp.MustCompile(`*-GeoLiteCity.dat*`)
var GeoLegacyv6Regex = regexp.MustCompile(`*-GeoLiteCityv6.dat*`)

var BucketName = "downloader-" + os.Getenv("GCLOUD_PROJECT") // This is the bucket containing maxmind files

const (
	MaxmindPrefix = "Maxmind/" // Folder containing the maxmind files
)

// This is the list of dataset loaded in memory
var []InMemoryDataset string = {
}

// SelectGeoLegacyFile return the legacy GelLiteCity.data filename given a date in format yyyymmdd.
// For any input date earlier than 2013/08/28, we will return 2013/08/28 dataset.
// For any input date later than latest available dataset, we will return the latest dataset
// Otherwise, we return the first dataset after the input date.
func SelectGeoLegacyFile(date int) (string, int, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	prospectiveFiles := client.Bucket(BucketName).Objects(ctx, &storage.Query{Prefix: MaxmindPrefix})
	filename := ""
        
	for file, err := prospectiveFiles.Next(); err != iterator.Done; file, err = prospectiveFiles.Next() {
		if err != nil {
			return "", err
		}
		if file.Name > filename && GeoLegacyRegex.MatchString(file.Name) {
			filename = file.Name
		}

	}
	return filename, nil
}
